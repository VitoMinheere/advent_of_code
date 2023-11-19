fn p1():
    let data = ""
    try:
        var f = open("test.txt", "r")
        let data = f.read()
        print(data)
        f.close()
    except Error:
        print("File not found")
    else:
        print(data)
    

fn main():
    p1()
