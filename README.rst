fileidentity
=================

This module provides feature to compare file identities.

.. warning::

   I test only on Mac OS X yet!

   I implement windows and plan9 based on golang libraries but they are not tested yet.

.. code-block:: bash

   $ go get github.com/shibukawa/fileidentity-go

API
------

* ``type fileidentity.FileIdentity struct``

  It stores OS specific data to describe file identity

* ``func (f *fileidentity.FileIdentity) Equals(o *fileidentity.FileIdentity) bool``

  Compare file identiteis.

* ``func fileidentity.NewFileIdentity(file *os.File) (*fileidentity.FileIdentity, error)``

  Read file identity from ``os.File``

* ``func fileidentity.NewFileIdentityFromFileInfo(file *os.FileInfo) (*fileidentity.FileIdentity, error)``

  Read file identity form ``os.FileInfo``. This is not avaiable on Windows.

License
---------

MIT

Author
---------

Yoshiki Shibukawa <yoshiki at shibu.jp>
