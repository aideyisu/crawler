import datetime

try:
    f = open('filelist.sh', 'r+')
    cyear = 2001
    cmonth = 10
    y, m = datetime.datetime.now().year,datetime.datetime.now().month
    # ./main http://archive.routeviews.org bgpdata 2001 10 RIBS
    while cyear != y or cmonth != m:
        if cmonth >= 10:
            f.write("./main http://archive.routeviews.org bgpdata " +
                    str(cyear) + " " + str(cmonth) + " RIBS\n")
            f.write("./main http://archive.routeviews.org bgpdata " +
                    str(cyear) + " " + str(cmonth) + " UPDATES\n")
        else:
            f.write("./main http://archive.routeviews.org bgpdata " +
                    str(cyear) + " 0" + str(cmonth) + " RIBS\n")
            f.write("./main http://archive.routeviews.org bgpdata " +
                    str(cyear) + " 0" + str(cmonth) + " UPDATES\n")
        if cmonth < 12:
            cmonth +=1
        elif cmonth == 12:
            cmonth = 1
            cyear += 1
finally:
    if f:
        f.close()
