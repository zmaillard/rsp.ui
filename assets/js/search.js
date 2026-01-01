import {instantMeiliSearch} from "@meilisearch/instant-meilisearch";
import instantsearch from "instantsearch.js";
import {searchBox, index, hits, configure} from "instantsearch.js/es/widgets";
import {connectInfiniteHits} from "instantsearch.js/es/connectors";


const SIGNBASEURL = document.getElementById('sign-base-url').value;
const SHIELDBASEURL = document.getElementById('shield-base-url').value;
const SEARCHINDEX = document.getElementById('search-index').value;
const HIGHWAYSEARCHINDEX = document.getElementById('highway-search-index').value;


const {searchClient} = instantMeiliSearch (
    document.getElementById('search-url').value,
    document.getElementById('search-api-key').value,
    {
        placeholderSearch: false, // default: true.
        primaryKey: 'id', // default: undefined
        meiliSearchParams: {
           showRankingScore: true,
        }
    },
);

let lastRenderArgs;


const infiniteHits = connectInfiniteHits(
    (renderArgs, isFirstRender) => {
        const {hits, showMore, widgetParams} = renderArgs;
        const { container } = widgetParams;

        lastRenderArgs = renderArgs;

        if (isFirstRender) {
            const sentinel = document.createElement('div');
            var $ul = document.createElement('ul');
            $ul.className = 'max-w-2xl divide-y divide-gray-200 dark:divide-gray-700';
            container.appendChild($ul);
            container.appendChild(sentinel);

            const observer = new IntersectionObserver(entries => {
                entries.forEach(entry => {
                    if (entry.isIntersecting && !lastRenderArgs.isLastPage) {
                        showMore();
                    }
                })
            });

            observer.observe(sentinel);

            return;
        }

        container.querySelector('ul').innerHTML = hits
            .map(
                hit =>
                    `<li class="pb-3 sm:pb-4">
                        <div class="flex items-center space-x-4">
                            <div class="shrink-0">
                                <a href="/sign/${hit.id}">
                                    <img class="w-32 h-32 rounded-sm" src="${SIGNBASEURL}${hit.id}/${hit.id}_t.jpg" alt="${hit.title}" />
                                </a>
                            </div>
                                     <div class="flex-1 min-w-0">
            <p class="text-sm font-medium text-gray-900 truncate dark:text-white">
                ${instantsearch.highlight({ attribute: 'title', hit })}
            </p>
            <p class="text-sm text-gray-500  dark:text-gray-400">
                ${instantsearch.highlight({ attribute: 'description', hit })}
            </p>
         </div>
                        </div>
                     </li>`
            ).join("");
    }
)

const search = instantsearch({
    indexName: SEARCHINDEX,
    searchClient,
    insights: false,
});

search.addWidgets([
    searchBox({
        container: '#searchbox',
        placeholder: 'Search for signs',
        showLoadingIndicator: true,
    }),
    infiniteHits({
        container: document.querySelector('#hits')
    }),
    index({indexName: HIGHWAYSEARCHINDEX}).addWidgets([
        configure({
            placeholderSearch: false, // default: true.
        }),
        hits({
            container: '#highway-hits',
            transformItems: items => items.filter(f=>f._rankingScore > 0.8),
            cssClasses: {
                list: ['max-w-2xl divide-y divide-gray-200 dark:divide-gray-700 result-list-item'],
                item: ['pb-3 sm:pb-4 result-item'],
            },
            templates: {
                empty: '', // Hide the empty message
                item: (hit, { html, components }) =>
                    html`
                        <div class="flex items-center space-x-4">
                            <div class="shrink-0">
                                <a href="/highway/${hit.slug}">
                                    <img class="w-32 h-32 rounded-sm" src="${SHIELDBASEURL}Shields/${hit.image_name}" alt="${hit.name}" />
                                </a>
                            </div>
                                     <div class="flex-1 min-w-0">
            <p class="text-sm font-medium text-gray-900 truncate dark:text-white">
                ${components.Highlight({ attribute: 'name', hit })}
            </p>
         </div>
                        </div>
                     `
            }
        })
    ]),
]);

search.start();