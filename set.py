import datetime


def main():
    try:
        f = open('filelist.sh', 'r+')
        cyear = 2001
        cmonth = 10
        y, m = datetime.datetime.now().year,datetime.datetime.now().month
        # ./main http://archive.routeviews.org bgpdata 2001 10 RIBS
        while cyear != y or cmonth != m:
            if m >= 10:
                f.write("./main http://archive.routeviews.org bgpdata " +
                        str(y) + " " + str(m) + " RIBS\n")
                f.write("./main http://archive.routeviews.org bgpdata " +
                        str(y) + " " + str(m) + " UPDATES\n")
            else:
                f.write("./main http://archive.routeviews.org bgpdata " +
                        str(y) + " 0" + str(m) + " RIBS\n")
                f.write("./main http://archive.routeviews.org bgpdata " +
                        str(y) + " 0" + str(m) + " UPDATES\n")
            print(y, m)
            if m <= 12 and m > 1:
                m -=1
            elif m == 1:
                m = 12
                y -= 1
    finally:
        if f:
            f.close()
    

if __name__ == '__main__':
    main()
