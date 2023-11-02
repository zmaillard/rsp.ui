import { autocomplete } from '@algolia/autocomplete-js';
import { instantMeiliSearch } from '@meilisearch/instant-meilisearch'
import instantsearch from 'instantsearch.js';
import historyRouter from 'instantsearch.js/es/lib/routers/history';
import { connectSearchBox } from 'instantsearch.js/es/connectors';
import {  hits, pagination } from 'instantsearch.js/es/widgets';



const searchClient = instantMeiliSearch (
    document.getElementById('search-url').value,
    document.getElementById('search-api-key').value,
    {
        placeholderSearch: false, // default: true.
        primaryKey: 'id', // default: undefined
    }
);

const INSTANT_SEARCH_INDEX_NAME = document.getElementById('search-index').value;
const SIGNBASEURL = document.getElementById('sign-base-url').value;

const instantSearchRouter = historyRouter();

const search = instantsearch({
    searchClient,
    indexName: INSTANT_SEARCH_INDEX_NAME,
    routing: instantSearchRouter,
});

// Mount a virtual search box to manipulate InstantSearch's `query` UI
// state parameter.
const virtualSearchBox = connectSearchBox(() => {});

search.addWidgets([
    virtualSearchBox({}),
    hits({
        container: '#hits',
        templates: {
            item(hit, { html, components }) {
            return html`<article class="media">
                    <figure class="media-left">
                        <p class="image is-4x3">
                            <a href="/sign/${hit.id}">
                                <img src="${SIGNBASEURL}${hit.id}/${hit.id}_t.jpg" />
                            </a>
                        </p>
                    </figure>
                    <div class="media-content">
                        <div class="content">
                            <p><strong><a href="/sign/${hit.id}">${components.Highlight({ attribute: 'title', hit })}</a></strong>
                            <br/>
                                ${hit.description}
                            </p>
                        </div>
                    </div>
        </article>`;
        },
    }}),
    pagination({
        container: '#pagination',
        templates: {

        }
    }),
]);

/*
<nav aria-label="Page navigation example">
  <ul class="flex items-center -space-x-px h-10 text-base">
    <li>
      <a href="#" class="flex items-center justify-center px-4 h-10 ml-0 leading-tight text-gray-500 bg-white border border-gray-300 rounded-l-lg hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">
        <span class="sr-only">Previous</span>
        <svg class="w-3 h-3" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 6 10">
          <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 1 1 5l4 4"/>
        </svg>
      </a>
    </li>
    <li>
      <a href="#" class="flex items-center justify-center px-4 h-10 leading-tight text-gray-500 bg-white border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">1</a>
    </li>
    <li>
      <a href="#" class="flex items-center justify-center px-4 h-10 leading-tight text-gray-500 bg-white border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">2</a>
    </li>
    <li>
      <a href="#" aria-current="page" class="z-10 flex items-center justify-center px-4 h-10 leading-tight text-blue-600 border border-blue-300 bg-blue-50 hover:bg-blue-100 hover:text-blue-700 dark:border-gray-700 dark:bg-gray-700 dark:text-white">3</a>
    </li>
    <li>
      <a href="#" class="flex items-center justify-center px-4 h-10 leading-tight text-gray-500 bg-white border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">4</a>
    </li>
    <li>
      <a href="#" class="flex items-center justify-center px-4 h-10 leading-tight text-gray-500 bg-white border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">5</a>
    </li>
    <li>
      <a href="#" class="flex items-center justify-center px-4 h-10 leading-tight text-gray-500 bg-white border border-gray-300 rounded-r-lg hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">
        <span class="sr-only">Next</span>
        <svg class="w-3 h-3" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 6 10">
          <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 9 4-4-4-4"/>
        </svg>
      </a>
    </li>
  </ul>
</nav>*/

search.start();

function navigate(id) {
    console.log(id)
    location.href = "/sign/" + id
}

// Set the InstantSearch index UI state from external events.
function setInstantSearchUiState(indexUiState) {
    search.setUiState(uiState => ({
        ...uiState,
        [INSTANT_SEARCH_INDEX_NAME]: {
            ...uiState[INSTANT_SEARCH_INDEX_NAME],
            // We reset the page when the search state changes.
            page: 1,
            ...indexUiState,
        },
    }));
}

// Return the InstantSearch index UI state.
function getInstantSearchUiState() {
    const uiState = instantSearchRouter.read();

    return (uiState && uiState[INSTANT_SEARCH_INDEX_NAME]) || {};
}

const searchPageState = getInstantSearchUiState();

let skipInstantSearchUiStateUpdate = false;
const { setQuery } = autocomplete({
    container: '#autocomplete',
    classNames: {
      input: 'block w-full p-4 pl-10 text-sm text-gray-900 border border-gray-300 rounded-lg bg-gray-50 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500',
    },
    placeholder: 'Search for signs',
    detachedMediaQuery: 'none',
    initialState: {
        query: searchPageState.query || '',
    },
    onSubmit({ state }) {
        setInstantSearchUiState({ query: state.query });
    },
    onReset() {
        setInstantSearchUiState({ query: '' });
    },
    onStateChange({ prevState, state }) {
        if (!skipInstantSearchUiStateUpdate && prevState.query !== state.query) {
            setInstantSearchUiState({ query: state.query });
        }
        skipInstantSearchUiStateUpdate = false;
    },
})

// This keeps Autocomplete aware of state changes coming from routing
// and updates its query accordingly
window.addEventListener('popstate', () => {
    skipInstantSearchUiStateUpdate = true;
    setQuery(search.helper?.state.query || '');
});