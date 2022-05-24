from flask import Flask, jsonify
import ray
from abc_live_lib.fetcher import Fetcher

app = Flask(__name__)

@app.route("/")
def news():
    fetcher = Fetcher()
    return jsonify({
        "data": fetcher.data
    }), 200

if __name__ == "__main__":
    ray.init(ignore_reinit_error = True)
    app.run()