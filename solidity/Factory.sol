// SPDX-License-Identifier: MIT
pragma solidity >=0.5.16 <0.9.0;
pragma experimental ABIEncoderV2;

import "./Storage.sol";

contract Factory {
    Storage private storageContract;

    event PrescriptionCreated(address prescriptionAddress);

    constructor(address _storageAddress) public {
        storageContract = Storage(_storageAddress);
    }

    function createPrescription(
        string memory _expirationDate,
        string memory _patient,
        string memory _dosageInstructions,
        string[] memory _medications,
        bool _isValid
    ) public {
        Prescription prescription = new Prescription(
            msg.sender,
            _expirationDate,
            _patient,
            _dosageInstructions,
            _medications,
            _isValid
        );

        storageContract.storePrescriptionAddress(address(prescription));
        emit PrescriptionCreated(address(prescription));
    }
}
