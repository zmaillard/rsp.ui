(ns sync-quality 
  (:require [core :as core]
            [pod.babashka.postgresql :as pg]
            [clojure.set :refer [join]])) 


(defn update-quality
  [imageid new-quality]
  (let [imageid-int (biginteger imageid)]
   (pg/execute! core/db ["UPDATE sign.highwaysign SET quality = ? WHERE imageid = ?" new-quality imageid-int])))

(defn -main 
  [& _args]
  (let [signs (map (fn[s] {:imageid (:annotation s) :eagle-star (:star s)})  (core/get-imported-signs))
        ext-signs (map (fn[s] {:imageid (:imageid s) :db-star (:vwindexsign/quality s) }) (core/get-signs))]
    (doseq [to-update (filter (fn[x] (not= (:eagle-star x) (:db-star x)))  (join signs ext-signs))]
      (update-quality (:imageid to-update) (:eagle-star to-update)))))