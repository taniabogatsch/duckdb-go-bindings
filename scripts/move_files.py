import os
import shutil
import argparse

if __name__ == "__main__":
    arg_parser = argparse.ArgumentParser(description='Moving half the files in a directory')
    arg_parser.add_argument(
        '-src', '--src', help='The source directory containing the files', required=True
    )
    arg_parser.add_argument(
        '-dst', '--dst', help='The destination directory containing the files', required=True
    )
    args = arg_parser.parse_args()

    files = os.listdir(args.src)
    for file in files[:len(files)//2]:
        shutil.move(args.src + file, args.dst + file)
