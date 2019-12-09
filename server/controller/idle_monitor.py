import psycopg2
import os
import json
import ast
import requests

def main():
    conn=psycopg2.connect(database="Tides",user="postgres",
            password="t6bB2T5KoQuPq6DrpWxJa3rYKVjIpOCtVSrKyBMB8PHcMShkidcQo8Kjn1lcXswB",host="10.11.16.83",port="30123")
    cur=conn.cursor()

    cur.execute('SELECT id, host_address, username, password, policy_id FROM resource_resource WHERE monitored = True')
    results = cur.fetchall()
    for result in results:
        i = result[0]
        #print(i)
        cur.execute('SELECT cpu, ram FROM usage_hostusage WHERE resource_id = %s', str(i))
        usage = cur.fetchone()

        print(usage)
        cur.execute('SELECT idle_policy, threshold_policy, template_id, is_destroy FROM policy_policy WHERE id = %s', str(result[4]))
        pols = cur.fetchone()
        deploy = False
        destroy = False
        #print(pols)
        ipol = eval(pols[0]) #idle policy
        bpol = eval(pols[1]) #busy policy
        
        if 'cpu' not in ipol.keys():
            if usage[1] < ipol['mem']:
                deploy = True
        elif 'mem' not in ipol.keys():
            if usage[0] < ipol['cpu']:
                deploy = True
        else:
            if usage[0] < ipol['cpu'] and usage[1] < ipol['mem']:
                deploy = True
                
        if 'cpu' not in bpol.keys():
            if usage[1] > bpol['mem']:
                destroy = True
        elif 'mem' not in bpol.keys():
            if usage[0] > bpol['cpu']:
                destroy = True
        else:
            if usage[0] > bpol['cpu'] or usage[1] > bpol['mem']:
                destroy = True


        if deploy:
            cur.execute("UPDATE resource_resource SET status = 'idle' WHERE id = %s", str(i))
            conn.commit()
            cur.execute("SELECT name FROM template_template WHERE id = %s", str(pols[2]))
            tname = cur.fetchone()
            #for name in tname:
                #file = os.path.join(settings.MEDIA_ROOT, 'uploads', name)
            os.system('python /home/shen1997/ve450/clone_vm.py -s ' + result[1] + ' -u ' + result[2] +' -p ' + result[3] +\
                ' --no-ssl --power-on --template ' + tname[0])

        elif destroy:
            cur.execute("UPDATE resource_resource SET status = 'busy' WHERE id = %s", str(i))
            conn.commit()
            if pols[3] is False:
                cur.execute("SELECT ip_address FROM usage_vmusage ORDER BY create_time DESC LIMIT 1")
                ip = cur.fetchone()
                data = {}
                data['ip_address'] = ip[0]
                requests.post("http://192.168.56.1:8000/api/usage/deletevm/", data=json.dumps(data))
                os.system('python /home/shen1997/ve450/destroy_vm.py -s ' + result[1] + ' -u ' + result[2] +' -p ' + result[3] +\
                    ' -i ' + ip[0])
            else:
                cur.execute("SELECT ip_address FROM usage_vmusage")
                ips = cur.fetchall()
                for ip in ips:
                    data = {}
                    data['ip_address'] = ip[0]
                    requests.post("http://192.168.56.1:8000/api/usage/deletevm/", data=json.dumps(data))
                    os.system('python /home/shen1997/ve450/destroy_vm.py -s ' + result[1] + ' -u ' + result[2] +' -p ' + result[3] +\
                    ' -i ' + ip[0])
                    
        else:
            cur.execute("UPDATE resource_resource SET status = 'normal' WHERE id = %s", str(i))
            conn.commit()


# start
if __name__ == "__main__":
    main()