(ns eagle-export-tags
  (:require [babashka.http-client :as http]
            [cheshire.core :as json]
            [clojure.java.io :as io]
            [clojure.data.csv :as csv]))
   

(def base-eagle-url (System/getenv "EAGLE_URL"))
(def eagle-token (System/getenv "EAGLE_TOKEN"))


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

(defn explode-tags
  [acc {imageid :imageid tags :tags}]
  (concat acc (map (fn[t][imageid t]) tags)))
 

(defn build-csv
  [signs]
  (->> signs
       (map (fn[s]{:imageid (:annotation s) :tags (:tags s)}))
       (reduce explode-tags [])))
    


(defn -main
  [& args]
  (when-let [[csv-path & _args] args]
   (let [signs (get-imported-signs)]
    (with-open [writer (io/writer csv-path)]
      (csv/write-csv writer (build-csv signs))))))
