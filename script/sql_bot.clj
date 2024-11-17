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

(defn build-db
  [path]
  (fs/delete-if-exists path)
  (sqlite/execute! path "create table sign (imageid TEXT, title TEXT, state TEXT, place TEXT, county TEXT, country TEXT)")
  (sqlite/execute! path "create table state (slug TEXT, name TEXT)")
  (sqlite/execute! path "create table country (slug TEXT, name TEXT)")) 
 

(defn load-sign
  [yaml-sign]
  (let [[imageid title state place county country] ((juxt :imageid :title :state :place :county :country) yaml-sign)]
    (sqlite/execute! db-path ["INSERT INTO sign (imageid, title, state, place, county, country) VALUES (?,?,?,?,?,?)" imageid title state place county country])))

(defn load-jurisdiction
  [table yaml-sign]
  (let [[slug name] ((juxt :slug :name) yaml-sign)
        insert (str "INSERT INTO " table  "(slug, name) VALUES (?,?)")]
    (sqlite/execute! db-path [insert slug name])))

(defn -main
  [& _args]
  (build-db db-path)
  (doseq [file (fs/glob "content/state" "**/*.md")]
   (load-jurisdiction "state" (yaml/parse-string (extract-yaml file))))
  (doseq [file (fs/glob "content/country" "**/*.md")]
   (load-jurisdiction "country" (yaml/parse-string (extract-yaml file))))
  (doseq [file (fs/glob "content/sign" "*.md")]
   (load-sign (yaml/parse-string (extract-yaml file)))))

  
