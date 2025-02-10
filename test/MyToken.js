const { expect } = require("chai");

describe("MyToken", function () {
  it("Should deploy and mint initial supply", async function () {
    const [owner] = await ethers.getSigners();

    const Token = await ethers.getContractFactory("MyToken");
    const token = await Token.deploy(owner.address);
    await token.deployed();

    expect(await token.balanceOf(owner.address)).to.equal(ethers.utils.parseUnits("1000000000", 18));
  });
});


