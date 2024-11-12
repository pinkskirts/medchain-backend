// SPDX-License-Identifier: MIT
pragma solidity >=0.5.16 <0.9.0;
pragma experimental ABIEncoderV2;

import "./Prescription.sol";

contract Storage {
    address[] public prescriptionAddresses;

    event AddressStored(address indexed prescriptionAddress);

    function storePrescriptionAddress(address _prescriptionAddress) public {
        prescriptionAddresses.push(_prescriptionAddress);
        emit AddressStored(_prescriptionAddress);
    }

    function getPrescriptionAddress(uint index) public view returns (address) {
        require(index < prescriptionAddresses.length, "Index out of bounds");
        return prescriptionAddresses[index];
    }

    function getAllPrescriptionAddresses() public view returns (address[] memory) {
        return prescriptionAddresses;
    }

    function getPrescriptionDetailsByAddress(address _prescriptionAddress) public view returns (
        address signingDoctor,
        string memory expirationDate,
        string memory patient,
        string memory dosageInstructions,
        string[] memory medications,
        bool isValid
    ) {
        Prescription prescription = Prescription(_prescriptionAddress);
        return prescription.getPrescriptionDetails();
    }
}
