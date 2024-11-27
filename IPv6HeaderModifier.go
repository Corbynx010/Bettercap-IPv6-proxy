package main

import (
        "github.com/chifflier/nfqueue-go/nfqueue"
        "fmt"
        "net"
)

func OnPacket(payload *nfqueue.Payload) int {
        fmt.Printf("Packet before modification:\n")
        listIPv6Headers(payload.Data)

        newPacket := editIPv6Header(payload.Data) // Modify rawPacket here

        // Re-parse the modified packet to ensure changes are reflected
        fmt.Printf("Packet after modification:\n")
        listIPv6Headers(newPacket)

        payload.SetVerdictModified(nfqueue.NF_ACCEPT, newPacket)

        return 0
}

func listIPv6Headers(packet []byte) {
        if len(packet) < 40 {
                fmt.Println("Packet too short to be an IPv6 packet")
                return
        }

        // IPv6 Header is always 40 bytes
        versionTrafficClassFlowLabel := packet[:4] // First 4 bytes contain version, traffic class, and flow label
        payloadLength := uint16(packet[4])<<8 | uint16(packet[5]) // Next 2 bytes
        nextHeader := packet[6]  // Next byte after payload length
        hopLimit := packet[7]   // Next byte after nextHeader
        sourceAddress := packet[8:24]  // Next 16 bytes for source address
        destAddress := packet[24:40]  // Next 16 bytes for destination address

        // Print out the IPv6 header fields
        fmt.Printf("IPv6 Header (First 40 bytes):\n")
        fmt.Printf("Version/Traffic Class/Flow Label: %x\n", versionTrafficClassFlowLabel)
        fmt.Printf("Payload Length: %d\n", payloadLength)
        fmt.Printf("Next Header: %d\n", nextHeader)
        fmt.Printf("Hop Limit: %d\n", hopLimit)
        fmt.Printf("Source Address: %s\n", net.IP(sourceAddress))
        fmt.Printf("Destination Address: %s\n", net.IP(destAddress))
        fmt.Printf("==========================================================\n")
}

func editIPv6Header(packet []byte) []byte {
        //versionTrafficClassFlowLabel := packet[:4] // First 4 bytes contain version, traffic class, and flow label
        copy(packet[:4], []byte{255, 255, 255, 255})
        //payloadLength := uint16(packet[4])<<8 | uint16(packet[5]) // Next 2 bytes
        //nextHeader := packet[6]  // Next byte after payload length
        //hopLimit := packet[7]   // Next byte after nextHeader
        //sourceAddress := packet[8:24]  // Next 16 bytes for source address
        //destAddress := packet[24:40]  // Next 16 bytes for destination address
        return packet
}