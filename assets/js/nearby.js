import {instantMeiliSearch} from "@meilisearch/instant-meilisearch";
import instantsearch from "instantsearch.js";
import { connectInfiniteHits} from "instantsearch.js/es/connectors";

const urlParams = new URLSearchParams(window.location.search);
const queryLat = urlParams.get('lat');
const queryLng = urlParams.get('lng');
const status = document.querySelector("#status");
const SIGNBASEURL = document.getElementById('sign-base-url').value;

const searchClient = instantMeiliSearch (
    document.getElementById('search-url').value,
    document.getElementById('search-api-key').value,
    {
        placeholderSearch: false, // default: true.
        primaryKey: 'id', // default: undefined
    },
);

const search = instantsearch({
    indexName: 'signs',
    searchClient,
    insights: false,
});

if (queryLat && queryLng) {
    const geoItem = {
        coords: {
            latitude: parseFloat(queryLat),
            longitude: parseFloat(queryLng)
        }
    }

    success(geoItem)
} else {

    if (!navigator.geolocation) {
        status.textContent = "Geolocation Not Supported By Your Browser";
    } else {
        status.textContent = "Locating…";
        navigator.geolocation.getCurrentPosition(success, error);
    }

}

function success(position) {
    const lat = position.coords.latitude;
    const lng = position.coords.longitude;

    status.textContent = ''
    const geoRadius = `_geoRadius(${lat},${lng},5000)`;
    search.addWidgets([{
        init: function(options) {
            options.helper.setQueryParameter('filters', geoRadius )
        }
    }]);

    search.start();
}

function error() {
    status.textContent = "Unable to retrieve your location";
}


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
                            <div class="flex-shrink-0">
                                <a href="/sign/${hit.id}">
                                    <img class="w-32 h-32 rounded" src="${SIGNBASEURL}${hit.id}/${hit.id}_t.jpg" alt="${hit.title}" />
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


search.addWidgets([
    infiniteHits({
        container: document.querySelector('#hits')
    })
]);
