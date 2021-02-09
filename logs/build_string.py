if __name__ == "__main__":
    with open("./logs/build_log.txt") as f:
        x = sum(1 for _ in f)
        print(f"{x:#08x}")