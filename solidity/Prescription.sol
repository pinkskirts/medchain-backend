// SPDX-License-Identifier: MIT
pragma solidity >=0.5.16 <0.9.0;
pragma experimental ABIEncoderV2;

contract Prescription {
    address public signingDoctor;
    string public expirationDate;
    string public patient;
    string public dosageInstructions;
    string[] public medications;
    bool public isValid;

    constructor(
        address _signingDoctor,
        string memory _expirationDate,
        string memory _patient,
        string memory _dosageInstructions,
        string[] memory _medications,
        bool _isValid
    ) public {
        signingDoctor = _signingDoctor;
        expirationDate = _expirationDate;
        patient = _patient;
        dosageInstructions = _dosageInstructions;
        medications = _medications;
        isValid = _isValid;
    }

    function getPrescriptionDetails() public view returns (
        address, string memory, string memory, string memory, string[] memory, bool
    ) {
        return (signingDoctor, expirationDate, patient, dosageInstructions, medications, isValid);
    }
}