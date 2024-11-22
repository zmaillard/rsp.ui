(ns sync-tags (:require
               [core :as core]
               [clojure.core :refer [biginteger]]
               [pod.babashka.postgresql :as pg]
               [clojure.set :refer [difference union join]]))

(def select-values (comp vals select-keys))

(defn create-tag
  [tag]
  (pg/execute! core/db ["INSERT INTO sign.tag (name, slug) VALUES (?, slugify(?))" tag tag]))


(defn delete-highway-tag
  [tagid imageid]
  (let [id (biginteger imageid)]
    (pg/execute! core/db ["DELETE FROM sign.tag_highwaysign WHERE tag_id = ? AND highwaysign_id IN (SELECT id from sign.highwaysign where imageid = ?)" tagid id])))

(defn insert-highway-tag
  [tagid imageid]
  (let [id (biginteger imageid)]
    (pg/execute! core/db ["INSERT INTO sign.tag_highwaysign (tag_id, highwaysign_id) VALUES (?, (SELECT id from sign.highwaysign where imageid = ?))" tagid id])))

(defn get-tag-ids
  [tag-sets]
  (let [db-tags (reduce #(union %1 %2) #{} (map #(:db-tags %1) tag-sets))
        eagle-tags (reduce #(union %1 %2)  #{} (map #(:eagle-tags %1) tag-sets))
        union-tags (union db-tags eagle-tags)
        db-res (pg/execute! core/db ["SELECT id, name FROM sign.tag WHERE name = any(?)" (into-array union-tags)])
        new-tags (difference union-tags (set (map #(:tag/name %1) db-res)))]
    (doseq[nt new-tags]
      (create-tag nt))
    (pg/execute! core/db ["SELECT id, name FROM sign.tag WHERE name = any(?)" (into-array union-tags)]))) 
       

(defn transform-tags
  [tags]
  (->> tags
   (map (fn[t] {(keyword(:tag/name t)) (:tag/id t)}))
   (reduce into {})))
       

(defn -main
  [& _args]
  (let [signs (map (fn [s] {:imageid (:annotation s) :eagle-tags (set (:tags s))})  (core/get-imported-signs))
        ext-signs (map (fn [s] {:imageid (:imageid s) :db-tags (set (:vwindexsign/tagitems s))}) (core/get-signs))
        changed (filter (fn [x] (seq (difference (:eagle-tags x) (:db-tags x))))  (join signs ext-signs))
        tag-ids (transform-tags(get-tag-ids changed))]
    (println tag-ids)
    (doseq [to-update changed]
      (let [image-id (:imageid to-update)
            to-delete (difference (:db-tags to-update) (:eagle-tags to-update))
            to-add (difference (:eagle-tags to-update) (:db-tags to-update))]
        (when-not (empty? to-delete) (doseq [tag to-delete] 
                                       (delete-highway-tag (get tag-ids (keyword tag)) image-id)))
        (when-not (empty? to-add) (doseq [tag to-add] 
                                    (insert-highway-tag (get tag-ids (keyword tag))image-id)))))))
      
