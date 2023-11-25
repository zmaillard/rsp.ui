FROM ghcr.io/osgeo/gdal:ubuntu-full-3.8.0 as gdal
ENV AWS_REGION=auto

WORKDIR /

COPY signs.json .

RUN ogr2ogr -f MBTILES /sign.mbtiles /signs.json -dsco MINZOOM=0 -dsco MAXZOOM=16

FROM ghcr.io/consbio/mbtileserver:0.10.0

COPY --from=gdal /sign.mbtiles /tilesets/sign.mbtiles
