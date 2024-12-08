#!/usr/bin/env bb
(ns eagle-import
  (:require [clojure.core :refer [biginteger]]
            [clojure.set :refer [join difference]]
            [babashka.http-client :as http]
            [cheshire.core :as json]
            [core :as core]
            [pod.babashka.postgresql :as pg]))


(def base-eagle-url (System/getenv "EAGLE_URL"))
(def eagle-token (System/getenv "EAGLE_TOKEN"))

(def db {:dbtype "postgresql"
         :host (System/getenv "DB_HOST")
         :dbname (System/getenv "DB_NAME")
         :user (System/getenv "DB_USER")
         :password (System/getenv "DB_PASSWORD")
         :port (System/getenv "DB_PORT")})



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


(defn create-folder
  [])

(defn transform-folder
  [folder-json-vec]
  (map #(hash-map :id (:id %) :name (:name %)) folder-json-vec))

(defn list-folders
  []
  (->
   (http/get (str base-eagle-url "/api/folder/list") {:query-params {:token eagle-token}})
   (get :body)
   (json/parse-string true)
   (get :data)
   (transform-folder)))



(defn get-signs
  []
  (pg/execute! db ["select imageid::text as imageid, title, state_name, tagitems, quality from sign.vwindexsign"]))

(defn find-new-signs
  [eagle-signs db-signs]

  (let [eagle-sign-ids (set (map :annotation eagle-signs))
        db-sign-ids (set (map :imageid db-signs))
        new-signs (difference db-sign-ids eagle-sign-ids)]
    (filter (fn [db-sign] (some #(= (:imageid db-sign)  %) new-signs))  db-signs)))

(defn build-path
  [sign]
  (let [base-path (System/getenv "BASE_LOCAL_IMAGE_PATH")
        image-id (:imageid sign)]

    (str base-path image-id  "/" image-id ".jpg")))


(defn build-website
  [sign]
  (let [image-id (:imageid sign)]
    (str "https://roadsign.pictures/sign/" image-id)))


(defn build-request
  [sign folders]
  (let [folderId (first (filter #(= (:name %) (:vwindexsign/state_name sign)) folders))]
    {:website (build-website sign)
     :path (build-path sign)
     :name (:vwindexsign/title sign)
     :tags (:vwindexsign/tagitems sign)
     :annotation (:imageid sign)
     :folderId (:id folderId)}))

(defn add-sign
  [sign folders]
  (let [req (json/encode (build-request sign folders))]
    (println req)
    (-> (http/post (str base-eagle-url "/api/item/addFromPath")
                  {:headers {:content-type "application/json"}
                   :body req
                   :query-params {:token eagle-token}})
        (get :body)
        (json/parse-string true)
        (get :status))))


(defn -main
  [& _args]
  (let [signs (find-new-signs (get-imported-signs) (get-signs))
        folders (list-folders)]
    (doseq [sign signs]
      (println (add-sign sign folders)))))
  
