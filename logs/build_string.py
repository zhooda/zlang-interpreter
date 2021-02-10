import sys
import subprocess
from datetime import date

args = sys.argv

if len(args) < 3:
    print(f"usage: {args[0]} <log_file> <go_file> <opt:semantic_version>")
    exit(-1)

log_file = args[1]
go_file = args[2]    
semantic_version = None
build_date = date.today().isoformat()
go_version = ' '.join(str(subprocess.check_output(["go", "version"]))[2:-3].split()[:-1])

if len(args) == 4:
    semantic_version = args[3]

if __name__ == "__main__":
    with open(log_file) as f:
        x = sum(1 for _ in f) + 1
        version_string = f"{x:#08x}"
        
    with open(go_file) as f:
        data = f.readlines()
        
    for i in range(len(data)):
        if 'buildString = "' in data[i]:
            data[i] = f'\tbuildString = "{version_string}"\n'
            
        if 'goVersion   = "' in data[i]:
            data[i] = f'\tgoVersion   = "{go_version}"\n'
            
        if 'buildDate   = "' in data[i]:
            data[i] = f'\tbuildDate   = "{build_date}"\n'
            
        if semantic_version:
            if 'version     = "' in data[i]:
                data[i] = f'\tversion     = "{semantic_version}"\n'
        
    with open(go_file, "w+") as f:
        f.writelines(data)
    
    print("Updated:")
    print(f"\tsemantic version: {semantic_version}")
    print(f"\tbuild string:     {version_string}")
    print(f"\tbuild date:       {build_date}")
    print(f"\tgo version:       {go_version}")
    # print(data)