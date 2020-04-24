import datetime

try:
    f = open('filelist.txt', 'r+')
    cyear = 2001
    cmonth = 1
    y, m = datetime.datetime.now().year,datetime.datetime.now().month
    # ./main http://archive.routeviews.org bgpdata 2001 10 RIBS
    while cyear != y or cmonth != m:
        f.write("./main http://archive.routeviews.org bgpdata " +
                str(cyear) + " " + str(cmonth) + " RIBS\n")
        f.write("./main http://archive.routeviews.org bgpdata " +
                str(cyear) + " " + str(cmonth) + " UPDATES\n")
        if cmonth < 12:
            cmonth +=1
        elif cmonth == 12:
            cmonth = 1
            cyear += 1
finally:
    if f:
        f.close()
