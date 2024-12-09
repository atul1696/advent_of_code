.SILENT: init
.PHONY: init
init:
	@echo "Enter Class name:"
	read name; \
	cp src/main/java/code/advent/Template.java src/main/java/code/advent/$$name.java; \
	sed -i '' "s#Template#$$name#g"  src/main/java/code/advent/$$name.java; \
	file_name=$$(echo $$name | tr '[:upper:]' '[:lower:]'); \
	touch "./src/main/resources/input/$$file_name.txt" "./src/main/resources/sample/$$file_name.txt";
