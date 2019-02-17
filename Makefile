target = http://localhost:8000
path = /hello
output = results.txt

run:
	echo "-------------------------------------------------------- ${path}" >> ${output}
	wrk -t12 -c400 -d1m --timeout 10s ${target}${path} >> ${output}
