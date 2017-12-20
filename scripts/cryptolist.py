import requests
import time
import datetime
import psycopg2
import psycopg2.extras
from settings import config

# from settings import DB_NAME, DB_USER, DB_HOST, DB_PASS
limit = "0"
url = "https://api.coinmarketcap.com/v1/ticker/?limit={0}".format(limit)
parsed_list = []

try:
    conn = psycopg2.connect(config['database']['uri'])
    cur = conn.cursor()
    print("Connection successful")
except:
    print("I am unable to connect to the database")

def get_list():
    r = requests.get(url)
    results = r.json()
    
    for result in results:
        if result['last_updated'] != None:
           last_updated = datetime.datetime.fromtimestamp(int(result['last_updated'])).strftime('%c')
            
        parsed_list.append((
            result['symbol'],
            result['name'],
            True,
            last_updated
        ))
    insert_query = "INSERT INTO cryptos (symbol, name, active, updated) VALUES %s ON CONFLICT DO NOTHING"
    psycopg2.extras.execute_values(cur, insert_query, parsed_list)
    conn.commit()

get_list()
