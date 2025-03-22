(ns core 
  (:require
   [babashka.http-client :as http]
   [cheshire.core :as json]
   [pod.babashka.postgresql :as pg]))


(def base-eagle-url (System/getenv "EAGLE_URL"))
(def eagle-token (System/getenv "EAGLE_TOKEN"))

(def db {:jdbcUrl (System/getenv "JDBC_DATABASE_URL")})

(defn update-quality
  [imageid new-quality]
  (let [imageid-int (biginteger imageid)]
   (pg/execute! db ["UPDATE sign.highwaysign SET quality = ? WHERE imageid = ?" new-quality imageid-int])))
  
(defn get-signs
  []
  (pg/execute! db ["select imageid::text as imageid, title, state_name, tagitems, quality from sign.vwindexsign"]))

(defn get-imported-signs
  []
  (->
   (loop [offset 0 signs []]
     (let [response (http/get (str base-eagle-url "/api/item/list") {:query-params {:limit 200 :offset offset :token eagle-token}})
           data (json/parse-string (:body response) true)
           new-signs (concat signs (:data data))]
       (if (empty? (:data data))
         new-signs
         (recur (+ offset 1) new-signs))))))

(defn get-tag-groups
  []
  (->
    (http/get (str base-eagle-url "/api/library/info"){:query-params {:token eagle-token}})
    (as-> r (json/parse-string (:body r) true))
    (get-in [:data :tagsGroups])))


