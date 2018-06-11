from flask import Flask
from flask import request
app = Flask(__name__)
@app.route('/login', methods=['POST', 'GET'])
def login():
    error = None
    if request.method == 'POST':
        if (request.form['username'] and request.form['password']):
            return ("username = "+request.form['username']+"\r\n")
        else:
            error = 'Invalid username/password'
    # the code below is executed if the request method
    # was GET or the credentials were invalid
    return render_template('login.html', error=error)
if __name__ == '__main__':
    app.run()
