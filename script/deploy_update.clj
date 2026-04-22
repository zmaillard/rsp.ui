(ns deploy-update
  (:require [babashka.curl :as curl]
           [cheshire.core :as json])) 


(def github-settings {:github-token (System/getenv "GITHUB_TOKEN")})

(defn -main 
 [& _args]
 (let [{token :github-token } github-settings
       url "https://api.github.com/repos/zmaillard/rsp.ui/actions/workflows/manual.yaml/dispatches"]
   (curl/post url
              {:headers {:accept "application/vnd.github+json"
                         "X-GitHub-Api-Version" "2026-03-10"
                         :authorization (str "Bearer " token)}
               :body (json/encode {:ref "main"})})))


