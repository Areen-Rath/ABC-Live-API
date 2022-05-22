from flask import Flask, jsonify
from mc_fetcher import mc_fetch
from et_fetcher import et_fetch
from bs_fetcher import bs_fetch

app = Flask(__name__)

@app.route("/")
def mc():
    data = mc_fetch()
    return jsonify({
        "data": data
    }), 200

@app.route("/economic_times")
def et():
    data = et_fetch()
    return jsonify({
        "data": data
    }), 200

@app.route("/business_standard")
def bs():
    data = bs_fetch()
    return jsonify({
        "data": data
    }), 200

if __name__ == "__main__":
    app.run(host = "0.0.0.0", port = 5000)