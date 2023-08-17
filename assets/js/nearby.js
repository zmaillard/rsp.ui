import { instantMeiliSearch } from '@meilisearch/instant-meilisearch'
import instantsearch from 'instantsearch.js';
import { connectSearchBox } from 'instantsearch.js/es/connectors';
import {  hits, pagination } from 'instantsearch.js/es/widgets';

const INSTANT_SEARCH_INDEX_NAME = document.getElementById('search-index').value;
const SIGN_BASE_URL = document.getElementById('sign-base-url').value;

const searchClient = instantMeiliSearch (
    document.getElementById('search-url').value,
    document.getElementById('search-api-key').value,
    {
        placeholderSearch: false, // default: true.
        primaryKey: 'id', // default: undefined
    }
);


const search = instantsearch({
    searchClient,
    indexName: INSTANT_SEARCH_INDEX_NAME,
});

const urlParams = new URLSearchParams(window.location.search);
const queryLat = urlParams.get('lat');
const queryLng = urlParams.get('lng');
const status = document.querySelector("#status");

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
                                <img src="${SIGN_BASE_URL}${hit.id}/${hit.id}_t.jpg" />
                            </a>
                        </p>
                    </figure>
                    <div class="media-content">
                        <div class="content">
                            <p><strong><a href="/sign/${hit.id}">${hit.title}</a></strong>
                            <br/>
                                ${hit.description}
                            </p>
                        </div>
                    </div>
        </article>`;
            },
        }})
]);


