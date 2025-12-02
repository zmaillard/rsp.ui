import * as params from '@params';

function loadWikiExcerpt(version) {
    let domWikiSearch = document.getElementById('wiki-search');
    if (!domWikiSearch) {
        return;
    }
    let externalPageLink =  domWikiSearch.dataset.externalLink;
    if (!externalPageLink) {
        return;
    }
    if (!("parse" in URL)) {
        return;
    }

    let url = URL.parse(externalPageLink);
    if (url.hash) {
        return;
    }
    let path = url.pathname;
    let title = path.substring(path.lastIndexOf('/') + 1);
    if (!title) return;
    let wikiApiUrl = `https://en.wikipedia.org/w/api.php?action=query&redirects=1&explaintext=true&exintro=true&prop=extracts&titles=${title}&format=json&origin=*`;

    fetch(wikiApiUrl, {method: "GET", headers: { "Api-User-Agent": `Roadsign Pictures/${version} (https://roadsign.pictures admin@roadsign.pictures)` }})
        .then(f=> f.json())
        .then(f=>{
            let keys = Object.keys(f.query.pages);
            if (keys.length > 0) {
                let pageId = keys[0];
                let text = f.query.pages[pageId].extract;

                if (text) {
                    domWikiSearch.innerHTML = `${text} Source: <a data-cy="wiki-source-link" class="hover:underline  inline-flex" rel="noopener noreferrer" target='_blank' href='${externalPageLink}'>Wikipedia</a>`;
                }
            }
        });
}

loadWikiExcerpt(params.version);
