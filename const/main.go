package main

import "fmt"

const x, y int = 1, 0x23

const s = "hello"

const c = "aa"

func main() {
	const d = 123

	const y = 123

	{
		const x = "aacd"
	}

	const (
		x1, y1      = 1, 2
		b      byte = byte(x)
		n           = uint8(12)
	)

	const (
		x2 = iota
		y2
		z2 = 100
		d1
		e = iota
		f
	)

	fmt.Println("iota",x2, y2, z2, d1, e, f)

	j := 2
	j = 3
	fmt.Println(&j)

	const k = 2
	b1 := k + 2
	fmt.Println(b1)

	type VersionKey int

	const (
		VersionBase VersionKey = iota   //0
		VersionRaftLogTruncationBelowRaft //1
		VersionSplitHardStateBelowRaft  //2
		VersionStatsBasedRebalancing  //3
		Version1_1                    //4
		VersionRaftLastIndex          //5
		VersionMVCCNetworkStats       //6
		VersionMeta2Splits            //7
		VersionRPCNetworkStats        //8
		VersionRPCVersionCheck        //9
		VersionClearRange             //10
		VersionPartitioning           //11
		VersionLeaseSequence          //12
		VersionUnreplicatedTombstoneKey  //13
		VersionRecomputeStats            //14
		VersionNoRaftProposalKeys        //15
		VersionTxnSpanRefresh            //16
		VersionReadUncommittedRangeLookups  //17
		VersionPerReplicaZoneConstraints    //18
		VersionLeasePreferences             //19

		// Add new versions here (step one of two).

	)

	fmt.Println("VersionBase:",VersionBase)
	fmt.Println("VersionRaftLogTruncationBelowRaft",VersionRaftLogTruncationBelowRaft)
	fmt.Println("VersionLeasePreferences",VersionLeasePreferences)


}
