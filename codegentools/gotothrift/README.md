This script will auto generate the go thrift client and server constructs for RPC.

There is a dependency on gobind and gotojson for this script.  Currently no checks are being made to ensure other scripts have been run.  The dependency resides in the fact that the gobind model should be generated as well as the JSON file.  Once the JSON file has been created if no deamons have registered for any objects then no thrift RPC will be generated.


