import psycopg2
from faker import Faker
from datetime import datetime, timedelta
from random import randint, choice, uniform

conn = None
try:
    conn = psycopg2.connect(
        host="localhost",
        database="uc",
        user="postgres",
        password="postgres")
except Exception:
    print("cannot connect")
    exit(-1)

cur = conn.cursor()
fake = Faker()


class User:
    def __init__(self, family, name, username, password, role) -> None:
        self.family = family
        self.name = name
        self.username = username
        self.password = password
        self.role = role

    def to_str(self):
        return f"{self.family}, {self.name}, {self.username}, {self.username}, {self.role}"

    def toTuple(self):
        return (self.family, self.name, self.username, self.username, self.role)


class Apartment:
    def __init__(self, size, room_numbers, userId):
        self.size = size
        self.roomNumbers = room_numbers
        self.userId = userId

    def toTuple(self):
        return (self.roomNumbers, self.size, self.userId)


class MonthlySpend:
    def __init__(self, _from, until, spend_amount, paid, apartment_id, service_id) -> None:
        self._from = _from
        self.until = until
        self.spend_amount = spend_amount
        self.paid = paid
        self.apartment_id = apartment_id
        self.service_id = service_id

    def toTuple(self):
        return (self._from,
                self.until,
                self.spend_amount,
                self.paid,
                self.apartment_id,
                self.service_id)


class ApartmentServices:
    def __init__(self, apartmentId, serviceId):
        self.apartmentId = apartmentId
        self.serviceId = serviceId

    def toTuple(self):
        return (self.apartmentId,
                self.serviceId)


def generateRandomDate():
    start_date = datetime.datetime(2020, 1, 1)
    end_date = datetime.datetime(2023, 5, 7)

    time_diff = (end_date - start_date).total_seconds()

    random_seconds = randint(0, int(time_diff))

    random_timestamp = start_date + datetime.timedelta(seconds=random_seconds)
    return random_timestamp


def generateUsers(usr_number):
    usrList = []
    for _ in range(usr_number):
        firstName = fake.first_name()
        lastName = fake.last_name()
        username = fake.last_name()
        password = username
        user_role = "user"
        user = User(firstName, lastName, username, password, user_role)
        usrList.append(user)
    return usrList


def generateApartments(apr_number):
    cur.execute("SELECT id FROM users")
    user_ids = [row[0] for row in cur.fetchall()]
    apartments = []
    sizes = [50, 60, 70, 80, 90, 100, 110, 120, 130, 140]
    room_numbers = [1, 2, 3, 4, 5]
    for i in range(apr_number):
        size = choice(sizes)
        room_number = choice(room_numbers)
        user_id = choice(user_ids)
        apartment = Apartment(size, room_number, user_id)
        apartments.append(apartment)
    return apartments

def generateBills(bill_number):
    aprList = []
    bills = []
    cur.execute("SELECT id FROM apartments")
    apartment_ids = [row[0] for row in cur.fetchall()]
    cur.execute("SELECT id FROM services")
    service_ids = [row[0] for row in cur.fetchall()]
    for i in range(bill_number):
        start_date = datetime(2022, 1, 1) + timedelta(days=randint(0, 365))
        end_date = start_date + timedelta(days=randint(1, 30))

        spend_amount = round(uniform(100, 1000), 2)
        paid = choice([True, False])
        apartment_id = choice(apartment_ids)
        service_id = choice(service_ids)

        monthly_spend = MonthlySpend(start_date, end_date, spend_amount, paid, apartment_id, service_id)
        bills.append(monthly_spend)
    return bills

usrNumber = 100
aprNumber = 66
billNumber = 200

def addUsers(usr_number):
    global conn, cur
    users = generateUsers(usr_number)
    for i in range(len(users)):
        postgres_insert_query = """insert into users(first_name, last_name, username, user_password, user_role) values
                                                                  (%s,%s,%s,%s,%s)"""
        user = users[i].toTuple()
        print(user)
        cur.execute(postgres_insert_query, user)
        conn.commit()

def addApartments(aprNumbers):
    global conn, cur
    apartments = generateApartments(aprNumbers)
    for i in range(len(apartments)):
        postgres_insert_query = """insert into apartments(room_numbers, size, user_id) values 
                                                          (%s,%s,%s)"""
        cur.execute(postgres_insert_query, apartments[i].toTuple())
        conn.commit()

def addBills(bill_number):
    global conn, cur
    bills = generateBills(bill_number)
    for i in range(len(bills)):
        postgres_insert_query = """insert into bills(\"from\", until, spend_amount, paid, apartment_id, service_id) values
                                                        (%s,%s,%s,%s,%s,%s)"""
        cur.execute(postgres_insert_query, bills[i].toTuple())
        conn.commit()

# addBills(billNumber)

conn.close()