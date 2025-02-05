package btree

import "encoding/binary"

const HEADER = 4
const BTREE_PAGE_SIZE = 4096 // Typical OS page size allocate 4kb
const BTREE_MAX_KEY_SIZE = 1000
const BTREE_MAX_VAL_SIZE = 3000

const BNODE_NODE = 1
const BNODE_LEAF = 2

type BNode []byte

type BTree struct {
	root 	uint64
	get 	func(uint64) []byte
	new		func([]byte) uint64
	del 	func(uint64)
}

// | type | nkeys | pointers |  offsets  | key-values | unused |
// |  2B  |   2B  | nkeys*8B |  nkeys*2B |    ...	  |  	   |

func (node BNode) btype() uint16 {
	return binary.LittleEndian.Uint16(node[0:2])
}

func (node BNode) nkeys() uint16 {
	return binary.LittleEndian.Uint16(node[2:4])
}

func (node BNode) setHeader(btype uint16, nkeys uint16) {
	binary.LittleEndian.PutUint16(node[0:2], btype)
	binary.LittleEndian.PutUint16(node[2:4], nkeys)
}

func (node BNode) getPtr(idx uint16) uint16 {
	pos := HEADER + 8*idx
	return binary.LittleEndian.Uint16(node[pos: ])
}

func (node BNode) setPtr(idx uint16, val uint64)

func offsetsPos(node BNode, idx uint16) uint16 {
	return HEADER + 8*node.nkeys() + 2*(idx-1)
}

func (node BNode) getOffset(idx uint16) uint16 {
	if idx == 0 {
		return 0
	}
	return binary.LittleEndian.Uint16(node[offsetsPos(node, idx):])
}

func (node BNode) setOffset(idx uint16, offset uint16)

func (node BNode) kvPos(idx uint16) uint16 {
	return HEADER + 8*node.nkeys() + 2*node.nkeys() + node.getOffset(idx)
}

func (node BNode) getKey(idx uint16) []byte {
	pos := node.kvPos(idx)
	klen := binary.LittleEndian.Uint16(node[pos:])
	return node[pos+4:][:klen]
}

func (node BNode) getVal(idx uint16) []byte

func (node BNode) nbytes() uint16 {
	return node.kvPos(node.nkeys())
}

func init() {
	node1max := HEADER + 8 + 2 + 4 + BTREE_MAX_KEY_SIZE + BTREE_MAX_VAL_SIZE
	if (node1max <= BTREE_PAGE_SIZE) {
		return
	}
}


