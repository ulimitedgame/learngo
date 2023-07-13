

CUR_DIR=$(shell pwd)
define remove_temp_files
	find $(CUR_DIR) -name "$1" -exec rm {} \;
endef

define list_files
	find $(CUR_DIR) -name "$1"
endef
temp_fils:= *.pyc *.pyo *.out 
info:
	@echo $(CUR_DIR)
	@echo $(temp_fils)
	$(foreach file,$(temp_fils), echo "find $(file)"; $(call list_files,$(file));)


clean:
	$(foreach file,$(temp_fils), echo "find $(file)"; $(call list_files,$(file));)
	@echo "\nremove temp files:"
	$(foreach file,$(temp_fils), echo "remove $(file)"; $(call remove_temp_files,$(file));)
	@echo "\nAfter remove temp files:"
	$(foreach file,$(temp_fils), echo "find $(file)"; $(call list_files,$(file));)
	


