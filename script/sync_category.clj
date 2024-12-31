(ns sync-category (:require
               [core :as core]
               [pod.babashka.postgresql :as pg]))


(defn set-categories
  [cats]
  (pg/execute! core/db ["UPDATE sign.tag SET is_category = false WHERE is_category = true"])
  (pg/execute! core/db ["UPDATE sign.tag SET is_category = true WHERE name = any(?)" (into-array cats)]))

(defn get-categories
  []
  (:tags (first(filter #(=(:name %1) "Categories") (core/get-tag-groups)))))

(defn -main
  []
  (set-categories(get-categories)))