; Calculate Low Qualtiy Image Placeholder for Image
; Adopted from article here: https://frzi.medium.com/lqip-css-73dc6dda2529
(ns lqip
  (:require [core :as core]
            [clojure.string :as str]
            [pod.babashka.postgresql :as pg]
            [babashka.process :refer [process]]))


(defn get-base64-thumbnail
 [path]
 (let [cmd (format "magick \"%s\" -resize 35x JPG:-", path)]
   (->
     (process cmd)
     (process {:out :string } "base64") deref :out)))

(defn build-base64-images 
  [i]
  (let [imageid (:annotation i)
        thumbnail-path (core/get-thumbnail-path (:id i))]
    {:imageId imageid :base64 (str/trim (get-base64-thumbnail thumbnail-path))}))


(defn update-thumbnail
  [{imageid :imageId thumbnail :base64}]
  (let [imageid-int (biginteger imageid)]
   (pg/execute! core/db ["UPDATE sign.highwaysign SET thumbnail = ? WHERE imageid = ?" thumbnail imageid-int])))

(defn get-existing-signs-missing-thumbnails
  []
  (pg/execute! core/db ["select imageid::text as imageid from sign.highwaysign where thumbnail is null"]))


(defn -main
 [& _args]
 (let 
  [ext (map #(:imageid %)  (get-existing-signs-missing-thumbnails))
   imp (core/get-imported-signs)]
  (doseq [s (filter (fn[c] (some (fn[a](= (:annotation c) a) )ext ))imp)]
       (update-thumbnail (build-base64-images s)))))
  
 
