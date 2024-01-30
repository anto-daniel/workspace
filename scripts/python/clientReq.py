# Python program to build a REST API using Flask

# Importing the required libraries
from flask import Flask, jsonify, request

# Creating an instance of Flask
app = Flask(__name__)


@app.route('/add', methods=['GET', 'POST'])
def home():
    if request.method == 'GET':
        data="Hello World"
        return jsonify({'data': data})


@app.route('/home/<int:num1>/<int:num2>', methods=['GET'])
def disp(num1, num2):
    if request.method == 'GET':
        data = num1 + num2
        return jsonify({'data': data})


if  __name__ == '__main__':
    app.run(debug=True)

