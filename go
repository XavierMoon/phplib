MACOSX_DEPLOYMENT_TARGET=10.5 ; \
GCC_ARCH="x86_64" ; \
CC=/usr/bin/gcc; \
CCFLAGS="-arch ${GCC_ARCH}" ; \
LDFLAGS="-Wl,-install_name,@rpath/libphp5.dylib -arch ${GCC_ARCH} -bind_at_load -lresolv" ; \
export CC LDFLAGS CFLAGS CCFLAGS MACOSX_DEPLOYMENT_TARGET ; \
./configure --build=${GCC_ARCH}-apple-darwin10.0.0  \
    --disable-debug --disable-rpath --disable-cli --enable-bcmath --enable-calendar \
    --enable-maintainer-zts --enable-embed=shared --enable-ftp \
    --enable-inline-optimization --enable-sockets --enable-wddx --sysconfdir=/etc/appweb \
    --with-pic --with-regex=system --with-pear --with-xmlrpc --with-zlib --with-curl ; \
cp Makefile Makefile.new ; \
sed 's/EXTRA_LDFLAGS = -avoid-version -module/EXTRA_LDFLAGS = -avoid-version/' < Makefile.new > Makefile ; \
make clean ; make

exit 0


./configure \
    --disable-debug --disable-rpath --disable-cli --enable-bcmath --enable-calendar \
    --enable-maintainer-zts --enable-embed=shared --enable-ftp \
    --enable-inline-optimization --enable-sockets --enable-wddx --sysconfdir=/etc/appweb \
    --with-pic --with-regex=system --with-pear --with-xmlrpc --with-zlib --with-curl ;
make clean
make
