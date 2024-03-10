# Vector Tile Generator

Converts a GeoJson file into a set of vector tiles using `ogr2ogr`.  Static tiles are hosted by mbtileserver.

<details>
  <summary>Dependencies</summary>
`signs.json` formatted as GeoJson.  Example:

```json
{
  "type": "FeatureCollection",
  "features": [
    {
      "type": "Feature",
      "geometry": {
        "type": "Point",
        "coordinates": [
          -123.031254,
          49.308176
        ]
      },
      "properties": {
        "imageid": "3487578858",
        "title": "TC-1 West - Exit 22"
      }
    },
    {
      "type": "Feature",
      "geometry": {
        "type": "Point",
        "coordinates": [
          -112.78828407097929,
          49.02530744674164
        ]
      },
      "properties": {
        "imageid": "1152953979",
        "title": "AB-62 North ABS-501 Jct."
      }
    },
    {
      "type": "Feature",
      "geometry": {
        "type": "Point",
        "coordinates": [
          -109.81734555919553,
          49.224232072708446
        ]
      },
      "properties": {
        "imageid": "1141974969",
        "title": "SK-13 East SK-21 South"
      }
    }
  ]
}
```
</details>

<details>
  <summary>Integrating Tiles Using Mapbox GL JS</summary>

```javascript
const map = new mapboxgl.Map({
    container: 'map', // container ID
    style: 'mapbox://styles/mapbox/streets-v12', // style URL
    center: [-95, 40], // starting position [lng, lat]
    zoom: 4 // starting zoom
});

map.on("load", () => {
    map.addSource('signs', {
        type: 'vector',
        tiles: ['https://yourmapserver.com/services/sign/tiles/{z}/{x}/{y}.pbf'],
        minzoom: 0,
        maxzoom: 16,
    });


    map.addLayer({
        id: "sign",
        type: "circle",
        source: "signs",
        "source-layer": "signs",
        paint: {
            "circle-radius": 8,
            "circle-color": "rgba(55,148,179,1)",
        },
    });
});
```
</details>