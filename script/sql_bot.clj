(ns sql-bot
  (:require [babashka.http-client :as http]
            [cheshire.core :as json]
            [babashka.fs :as fs]
            [clojure.java.io :as io]
            [clojure.string :refer (join)]
            [clj-yaml.core :as yaml]
            [pod.babashka.go-sqlite3 :as sqlite]))

(def db-path "bot.db")

(defn extract-yaml
  [file]
  (with-open [rdr (io/reader (str file))]
    (let [[_front-matter & line] (line-seq rdr)]
     (->> line
          (take-while #(not= %1 "---"))
          (join "\n")))))

(defn long-string
  [& strings]
  (join "\n" strings))

(defn sql-view
  []
  (long-string
     "CREATE VIEW vwSign AS"
     "SELECT s.imageid, s.title, st.name as state, pl.name as place, ct.name as county, cn.name as country, s.quality"
     "FROM main.sign as s"
     "INNER JOIN main.country as cn ON s.country = cn.slug"
     "INNER JOIN main.state as st ON s.state = st.slug"
     "LEFT OUTER JOIN main.county as ct ON s.county = ct.key"
     "LEFT OUTER JOIN main.place as pl ON s.place = pl.key" ))


(defn build-db
  [path]
  (fs/delete-if-exists path)
  (sqlite/execute! path "create table sign (imageid TEXT, title TEXT, state TEXT, place TEXT, county TEXT, country TEXT, quality INTEGER)")
  (sqlite/execute! path "create table state (slug TEXT PRIMARY KEY, name TEXT)")
  (sqlite/execute! path "create table place (slug TEXT, name TEXT, state TEXT, key TEXT PRIMARY KEY)")
  (sqlite/execute! path "create table county (slug TEXT, name TEXT, state TEXT, key TEXT PRIMARY KEY)")
  (sqlite/execute! path "create table country (slug TEXT PRIMARY KEY, name TEXT)")
  (sqlite/execute! path (sql-view)))

(defn load-sign
  [yaml-sign]
  (let [[imageid title state place county country quality] ((juxt :imageid :title :state :place :county :country :quality) yaml-sign)]
    (sqlite/execute! db-path ["INSERT INTO sign (imageid, title, state, place, county, country, quality) VALUES (?, ?,?,?,?,?,?)" imageid title state place county country quality])))

(defn load-jurisdiction
  [table yaml-sign]
  (let [[slug name] ((juxt :slug :name) yaml-sign)
        insert (str "INSERT INTO " table  "(slug, name) VALUES (?,?)")]
    (sqlite/execute! db-path [insert slug name])))

(defn load-jurisdiction-with-state
  [table yaml-sign]
  (let [[slug name state] ((juxt :slug :name :stateslug) yaml-sign)
        insert (str "INSERT INTO " table  "(slug, name, state, key) VALUES (?,?,?,?)")]
    (sqlite/execute! db-path [insert slug name state (str state "_" slug)])))

(defn -main
  [& _args]
  (build-db db-path)
  (doseq [file (fs/glob "content/county" "**/*.md")]
   (load-jurisdiction-with-state "county" (yaml/parse-string (extract-yaml file))))
  (doseq [file (fs/glob "content/place" "**/*.md")]
   (load-jurisdiction-with-state "place" (yaml/parse-string (extract-yaml file))))
  (doseq [file (fs/glob "content/state" "**/*.md")]
   (load-jurisdiction "state" (yaml/parse-string (extract-yaml file))))
  (doseq [file (fs/glob "content/country" "**/*.md")]
   (load-jurisdiction "country" (yaml/parse-string (extract-yaml file))))
  (doseq [file (fs/glob "content/sign" "*.md")]
   (load-sign (yaml/parse-string (extract-yaml file)))))

  
