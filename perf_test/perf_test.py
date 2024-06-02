import paramiko
import getpass
from faker import Faker
import random

total_test_insert = 100000 # определяем, сколько adverts, likes и users будем создавать

def execute_remote_command(ssh_client, command):
    stdin, stdout, stderr = ssh_client.exec_command(command)
    stdout.channel.recv_exit_status()
    return stdout.read().decode(), stderr.read().decode()


def recreate_tables(ssh_client):
    target_directory = '/home/root/2024_1_TeaStealers/'

    psql_command = f"cd {target_directory} && echo 'y' | make migrate-down"
    result, error = execute_remote_command(ssh_client, psql_command)

    psql_command = f"cd {target_directory} && make migrate-up"
    result, error = execute_remote_command(ssh_client, psql_command)


def generate_random_number_string():
    length = random.randint(8, 15)
    random_number_string = ''.join(random.choices('0123456789', k=length))
    return random_number_string

def generate_random_email():
    length = random.randint(8, 30)
    random_number_string = ''.join(random.choices('ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz@.', k=length))
    return random_number_string

def generate_random_password():
    length = random.randint(9, 38)
    random_number_string = ''.join(random.choices('ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789', k=length))
    return random_number_string


def insert_user_data(ssh_client):
    fake = Faker()
    container_name = 'postgres'
    database_name = 'tean'
    database_user = 'tean_user'
    i = 0 # счётчик

    insert_query = (f"INSERT INTO user_data (password_hash, level_update, phone, email, created_at) VALUES ")

    limit = 10
    for _ in range(limit):
        password_hash = generate_random_password()
        level_update = fake.random_int(min=1, max=100)
        phone = generate_random_number_string()
        email = generate_random_email()
        created_at = fake.date_time_this_decade().isoformat()


        insert_query +=  f"('{password_hash}', {level_update}, '{phone}', '{email}', '{created_at}')"
        if i == limit-1:
            insert_query += ";"
        else:
            insert_query += ", "
        i+=1

    psql_command = (
        f"docker exec -i {container_name} psql -U {database_user} -d {database_name} -c \"{insert_query}\""
    )

    result, error = execute_remote_command(ssh_client, psql_command)
    if error:
        print(f"Error executing insert query: {error}")


def insert_advert_data(ssh_client):
    fake = Faker()
    container_name = 'postgres'
    database_name = 'tean'
    database_user = 'tean_user'
    i = 0

    insert_query = (
        "INSERT INTO advert "
        "(user_id, title, type_placement, description, phone, is_agent, priority, created_at, is_deleted) "
        "VALUES "
    )

    limit = 10
    for _ in range(limit):
        user_id = fake.random_int(min=1, max=total_test_insert)
        title = fake.text(max_nb_chars=70)
        type_placement = fake.random_element(elements=('Rent', 'Sale'))
        description = fake.text(max_nb_chars=400)
        phone = generate_random_number_string()
        is_agent = fake.boolean()
        priority = fake.random_int(min=1, max=10000)
        created_at = fake.date_time_this_decade().isoformat()

        insert_query += (
            f"({user_id}, '{title}', '{type_placement}', '{description}', '{phone}', {is_agent}, {priority}, "
            f"'{created_at}', FALSE)"
        )
        if i == limit-1:
            insert_query += ";"
        else:
            insert_query += ", "
        i += 1

    psql_command = (
        f"docker exec -i {container_name} psql -U {database_user} -d {database_name} -c \"{insert_query}\""
    )

    result, error = execute_remote_command(ssh_client, psql_command)
    if error:
        print(f"Error executing insert query: {error}")


def insert_favourite_advert_data(ssh_client):
    fake = Faker()
    container_name = 'postgres'
    database_name = 'tean'
    database_user = 'tean_user'
    i = 0

    insert_query = (
        "INSERT INTO favourite_advert "
        "(user_id, advert_id, is_deleted) "
        "VALUES "
    )

    limit = 10
    for _ in range(limit):
        user_id = fake.random_int(min=1, max=total_test_insert)
        advert_id = fake.random_int(min=1, max=total_test_insert)
        is_deleted = False

        insert_query += (
            f"({user_id}, {advert_id}, {is_deleted})"
        )
        if i == limit-1:
            insert_query += ";"
        else:
            insert_query += ", "
        i += 1

    psql_command = (
        f"docker exec -i {container_name} psql -U {database_user} -d {database_name} -c \"{insert_query}\""
    )

    result, error = execute_remote_command(ssh_client, psql_command)
    if error:
        print(f"Error executing insert query: {error}")

def main():
    ssh_host = '78.155.197.35'
    ssh_port = 22
    ssh_user = 'root'
    ssh_password = getpass.getpass(prompt="Enter SSH password: ")

    ssh_client = paramiko.SSHClient()
    ssh_client.set_missing_host_key_policy(paramiko.AutoAddPolicy())

    try:
        ssh_client.connect(
            hostname=ssh_host,
            port=ssh_port,
            username=ssh_user,
            password=ssh_password
        )

        if ssh_client is None:
            print("No connection")
            exit(-1)

        recreate_tables(ssh_client)

        count = int(total_test_insert/10)
        for _ in range(count):
            insert_user_data(ssh_client)

        for _ in range(count):
            insert_advert_data(ssh_client)

        for _ in range(count):
            insert_favourite_advert_data(ssh_client)

    finally:
        ssh_client.close()


def main_bench():
    ssh_host = '78.155.197.35'
    ssh_port = 22
    ssh_user = 'root'
    ssh_password = getpass.getpass(prompt="Enter SSH password: ")

    ssh_client = paramiko.SSHClient()
    ssh_client.set_missing_host_key_policy(paramiko.AutoAddPolicy())

    try:
        ssh_client.connect(
            hostname=ssh_host,
            port=ssh_port,
            username=ssh_user,
            password=ssh_password
        )

        if ssh_client is None:
            print("No connection")
            exit(-1)


        container_name = 'postgres'
        database_name = 'tean'
        database_user = 'tean_user'

        query = "SELECT * FROM advert WHERE id = 1400;"
        psql_command = (
            f"docker exec -i {container_name} psql -U {database_user} -d {database_name} -c \"{query}\""
        )

        query_result, error = execute_remote_command(ssh_client, psql_command)

        if error:
            print(f"Error executing query: {error}")
            exit(-1)
        else:
            print(f"Query result:")
            print(query_result)



        id = random.randint(1, total_test_insert)
        psql_command = (
            f"cd /home/root/wrk-test/wrk && wrk -t4 -c100 -d30s http://tean.homes/api/test/count/{id}"
        )

        result, error = execute_remote_command(ssh_client, psql_command)

        if error:
            print(f"Error doing test: {error}")
            exit(-1)
        else:
            print(f"Test Result count likes:")
            print(result)

        psql_command = (
            f"cd /home/root/wrk-test/wrk && wrk -t4 -c100 -d30s http://tean.homes/api/test/fast/{id}"
        )

        result, error = execute_remote_command(ssh_client, psql_command)

        if error:
            print(f"Error doing test: {error}")
            exit(-1)
        else:
            print(f"Test Result fast get:")
            print(result)

    finally:
        ssh_client.close()


if __name__ == "__main__":
    # main()
    main_bench()