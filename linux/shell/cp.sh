printfile()
{
    srcdir=$1
    destdir=$2
    for file in "$srcdir"/*
    do
        if [ -d $file ]
        then
            printfile $file $destdir
        else
            cp $file $destdir
        fi
    done
}
 
printfile $(pwd)/$1 $(pwd)/$2

# cp.sh source/ dest/