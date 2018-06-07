import os
import sys
import glob
import csv
from xlsxwriter.workbook import Workbook


def main():
    if len(sys.argv) < 2:
        print('./bin [input dir] [output dir]')
        exit(1)

    input_dir = sys.argv[1]
    output_dir = sys.argv[2]

    if not os.path.exists(input_dir) or not os.path.isdir(input_dir):
        print('{} is invalid'.format(input_dir))
        exit()
    if os.path.exists(output_dir) and not os.path.isdir(output_dir):
        print('{} is invalid'.format(output_dir))
        exit()
    if input_dir == output_dir:
        print('input/output directory should not equal')
        exit()
    if not os.path.exists(output_dir):
        os.mkdir(output_dir)

    for path, dirs, files in os.walk(input_dir):
        #print(path, dirs, files)
        for f in files:
            print(path)
            if not os.path.exists('{}/{}'.format(output_dir, path)):
                print('{}/{}'.format(output_dir, path))
                os.makedirs('{}/{}'.format(output_dir, path))
            store_path = '{}/{}/{}.xlsx'.format(output_dir, path, f[:-4])
            print(store_path)
            read_path = '{}/{}'.format(path, f)
            workbook = Workbook(store_path, {'strings_to_numbers': True})
            worksheet = workbook.add_worksheet()
            with open(read_path, 'rt', encoding='utf8') as f:
                reader = csv.reader(f)
                for r, row in enumerate(reader):
                    for c, col in enumerate(row):
                        worksheet.write(r, c, col)
            workbook.close()
if __name__ == '__main__':
    main()
