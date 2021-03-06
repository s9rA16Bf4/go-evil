from os import path
from subprocess import check_output, run
from re import search

def get_current_compiler_version() -> str:
    """ This function will run gevil passing the -v flag, and regex the result """
    output = str(check_output(["./gevil", "-v"]))
    regex = search("([0-9]\\.[0-9])", output)
    return regex.group()

def main():
    current_version = get_current_compiler_version()
    for examples in open("examples/examples.txt", "r"): # Get every found example
        examples = examples.rstrip("\n")
        with open(examples, "r") as out:
            content = out.readlines()
        
        for i, line in enumerate(content):
            if "[version " in line:
                script_version = line.split(" ")[1]
                script_version = script_version.replace("]  ", "")       
                if script_version != current_version: # We need to update
                    content[i] = f"[version {current_version}]\n" # Simply replaces this line 

        with open(examples, "w") as out:
            out.writelines(content) # Then we overwrite the file but with our newly modified section

if __name__ == "__main__":
    if not path.exists("examples/examples.txt"):
        run(["python3", "examples/generate_list_of_examples.py"]) # Needs to exist

    main()