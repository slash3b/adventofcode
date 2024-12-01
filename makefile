curr_year:=$(shell date +%Y)
workdir:=$(shell pwd)

day: year
	@cp day.tpl ${curr_year}/

year:
	@mkdir -p ${workdir}/${curr_year}
	@mkdir -p ${workdir}/${curr_year}/inputs
