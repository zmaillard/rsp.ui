import { autocomplete } from '@algolia/autocomplete-js';
import algoliasearch from 'algoliasearch/lite';
import { instantMeiliSearch } from '@meilisearch/instant-meilisearch'
import instantsearch from 'instantsearch.js';
import historyRouter from 'instantsearch.js/es/lib/routers/history';
import { connectSearchBox } from 'instantsearch.js/es/connectors';
import { hierarchicalMenu, hits, pagination } from 'instantsearch.js/es/widgets';



const searchClient = instantMeiliSearch (
    document.getElementById('search-url').value,
    document.getElementById('search-api-key').value,
    {
        placeholderSearch: false, // default: true.
        primaryKey: 'id', // default: undefined
    }
);

const INSTANT_SEARCH_INDEX_NAME = document.getElementById('search-index').value;
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
                                <img src="https://sign.sagebrushgis.com/${hit.id}/${hit.id}_t.jpg" />
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
    }),
]);

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