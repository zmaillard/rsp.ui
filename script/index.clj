(ns index
  (:require [babashka.http-client :as http]
            [cheshire.core :as json]
            [pod.babashka.postgresql :as pg]))

(def batch-size 25)

(def db {:jdbcUrl (System/getenv "JDBC_DATABASE_URL")})

(def index-settings {:host (System/getenv "HUGO_PARAMS_SEARCHURL")
                     :sign-index (System/getenv "HUGO_PARAMS_SEARCHINDEX")
                     :highway-index (System/getenv "HUGO_PARAMS_SEARCHINDEXHIGHWAY")
                     :key (System/getenv "HUGO_PARAMS_SEARCHKEY")})

(defn update-last-updated
  []
  (pg/execute! db ["UPDATE sign.highwaysign SET last_indexed = now() WHERE last_indexed is null or last_indexed < last_update"]))

(defn build-image-url
  [sign]
  (let [image-id (:imageid sign)
        hosting-url (System/getenv "HUGO_PARAMS_SIGNBASEURL")]
      (str hosting-url image-id "/" image-id "_l.jpg")))



(defn build-locale
  [name slug]
  (if name {:name name :slug slug} nil)) 

(defn build-sign-request
  [sign]
  {:id (:imageid sign)
   :title (:vwindexsign/title sign)
   :description (:vwindexsign/sign_description sign)
   :_geo {:lng (:longitude sign) :lat (:latitude sign)}
   :date_taken (:vwindexsign/date_taken sign)
   :highways (map (fn[h] {:name (:name h) :slug (:slug h)}) (:vwindexsign/hwys sign))
   :country (build-locale (:vwindexsign/country_name sign) (:vwindexsign/country_slug sign))  
   :county (build-locale (:vwindexsign/county_name sign) (:vwindexsign/county_slug sign))  
   :state (build-locale (:vwindexsign/state_name sign) (:vwindexsign/state_slug sign))   
   :url (build-image-url sign) 
   :quality (:vwindexsign/quality sign)
   :place (build-locale (:vwindexsign/place_name sign) (:vwindexsign/place_slug sign))})

(defn build-highway-request
  [hwy]
  {:id (:highway/id hwy)
   :name (:highway/highway_name hwy)
   :slug (:highway/slug hwy)
   :image_name (:highway/image_name hwy)
   :url (:case hwy)})


(defn get-signs
  []
  (let [res (pg/execute! db ["SELECT imageid::text, title, sign_description, date_taken, country_slug, country_name, state_slug, state_name, county_name, county_slug, place_name, place_slug, hwys, ST_X(point::geometry) as longitude, ST_Y(point::geometry) as latitude, quality FROM sign.vwindexsign where last_indexed is null or last_indexed < last_update"])] 
   (map build-sign-request res)))

(defn get-highways
  []
  (let [res (pg/execute! db ["select id, highway_name, image_name, slug, case when image_name = '' then '' when image_name is null then '' else 'https://shield.roadsign.pictures/Shields/' || highway.image_name end from sign.highway where id in (select distinct highway_id from sign.highwaysign_highway)"])] 
   (map build-highway-request res)))


(defn update-sign-index
  [signs]
  (let [req (json/encode signs)
        {host :host index :sign-index key :key} index-settings]
    (-> (http/put (str host "/indexes/" index "/documents")
              {:headers {:content-type "application/json"
                              :authorization (str "Bearer " key)}
                    :body req})
        (get :status))))
     

(defn update-highway-index
  [signs]
  (let [req (json/encode signs)
        {host :host index :highway-index key :key} index-settings]
    (-> (http/put (str host "/indexes/" index "/documents")
                  {:headers {:content-type "application/json"
                             :authorization (str "Bearer " key)}
                   :body req})
        (get :status))))
     


(defn apply-updates
  [coll fn]
  (doseq [coll-block (partition-all batch-size coll)]
    (println(fn coll-block))))
  

(defn -main
  [& _args]
  (apply-updates (get-signs) update-sign-index)
  (apply-updates (get-highways) update-highway-index)
  (update-last-updated))
