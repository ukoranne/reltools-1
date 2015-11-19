The GO bind utilizes work from pyangbind

Pre-requesitis are Python 2.7 and pyang 1.6 (may work with other version but work was tested against these versions)

to execute here is an example:
cd ~/git/external/src/github.com/openconfig/release/models/bgp
// pick a file to run against I chose openconfig-bgp.yang 
pyang ~/git/reltools/codegentools/gobind -f pybind openconfig-bgp.yang -o test.go	

code expects the -o test.go and it will convert it to fmt_test.go. 



11/19/15 - Created another gobind which will collapse the leaf members down to the last child in the tree and add the parent leafs/leaf-lists to the struct 

code was run against openconfig interface
pyang --plugindir ~/git/reltools/codegentools/gobind -f pybind -o if.go *.yang



