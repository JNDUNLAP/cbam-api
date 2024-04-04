from gen.create.xsd_processor import process_xsd_files
from gen.create.go_struct_generator import generate_go_struct
import json
import shutil
import os


json_output_file = 'req/element_details.json'
golang_output_file = 'report_struct.go'
xsd_files = ['req/QReport_17.03.xsd', 'req/stypes_17.03.xsd']
target_path = '../app/model/report_struct.go'
target_dir = os.path.dirname(target_path)



def main():

    element_details = process_xsd_files(xsd_files)

    with open(json_output_file, 'w') as json_file:
        json.dump(element_details, json_file, indent=4)

    generate_go_struct(json_output_file, golang_output_file)
    shutil.move(golang_output_file, target_path)

    print("Process completed. The Go struct file has been generated.")

if __name__ == "__main__":
    main()
