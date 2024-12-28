(ns sync-category (:require
               [core :as core]
               [clojure.core :refer [biginteger]]
               [pod.babashka.postgresql :as pg]
               [clojure.set :refer [difference union join]]))

(defn -main
  []
  (println(:tags (first(filter #(=(:name %1) "Categories") (core/get-tag-groups))))))