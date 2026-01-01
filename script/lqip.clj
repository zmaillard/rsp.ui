(ns lqip
  (:require [clojure.math :refer [round]]
            [core :as core]
            [pod.babashka.postgresql :as pg]
            [babashka.http-client :as http]
            [cheshire.core :as json]))


(def base-eagle-url (System/getenv "EAGLE_URL"))
(def eagle-token (System/getenv "EAGLE_TOKEN"))


(defrecord RGB [r g b])


(defn build-color [{[r g b] :color}]
  (RGB. r g b))

(defn build-palette [i]
  {:imageId (:annotation i) :c0  (build-color (first(:palettes i))) :c1 (build-color (second(:palettes i))) :c2 (build-color (nth(:palettes i) 2))})


(defn pack-color-10-bit [^RGB c]
  (let [r (round(* (/ (:r c) 0xFF) 2r1111))
        g (round(* (/ (:g c) 0xFF) 2r1111))
        b (round(* (/ (:b c) 0xFF) 2r111))]
       (bit-or (bit-shift-left r 7) (bit-shift-left g 3) b)))

(defn pack-color-11-bit [^RGB c]
  (let [r (round(* (/ (:r c) 0xFF) 2r111))
        g (round(* (/ (:g c) 0xFF) 2r1111))
        b (round(* (/ (:b c) 0xFF) 2r111))]
       (bit-or (bit-shift-left r 7) (bit-shift-left g 3) b)))


(defn combine-colors [^RGB c0 ^RGB c1 ^RGB c2]
  (let [pc0 (pack-color-11-bit c0)
        pc1 (pack-color-11-bit c1)
        pc2 (pack-color-10-bit c2)
        combined (bit-or (bit-shift-left pc0 21)(bit-shift-left pc1 10) pc2)]
   (format "#%08x" combined)))

(defn update-lqip
  [imageid new-lqip]
  (let [imageid-int (biginteger imageid)]
   (pg/execute! core/db ["UPDATE sign.highwaysign SET lqip_hash = ? WHERE imageid = ?" new-lqip imageid-int])))


(defn -main
  [& args]
  (let [signs (map (fn[s] {:imageId (:annotation s) :palette (build-palette s)})(core/get-imported-signs))]
    (doseq [{i :imageId c :palette} signs] 
     (update-lqip i (combine-colors (:c0 c)(:c1 c)(:c2 c))))))

;(defn -main 
;  [& _args]
;  (let [signs (map (fn[s] {:imageid (:annotation s) :eagle-star (:star s)})  (core/get-imported-signs))
;        ext-signs (map (fn[s] {:imageid (:imageid s) :db-star (:vwindexsign/quality s) }) (core/get-signs))]
;    (doseq [to-update (filter (fn[x] (not= (:eagle-star x) (:db-star x)))  (join signs ext-signs))]
;      (update-quality (:imageid to-update) (:eagle-star to-update)))))
  
