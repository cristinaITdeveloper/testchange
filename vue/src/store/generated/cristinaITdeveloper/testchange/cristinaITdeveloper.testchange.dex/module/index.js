// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgSendSellOrder } from "./types/dex/tx";
import { MsgCancelSellOrder } from "./types/dex/tx";
import { MsgSendCreatePair } from "./types/dex/tx";
import { MsgCancelBuyOrder } from "./types/dex/tx";
const types = [
    ["/cristinaITdeveloper.testchange.dex.MsgSendSellOrder", MsgSendSellOrder],
    ["/cristinaITdeveloper.testchange.dex.MsgCancelSellOrder", MsgCancelSellOrder],
    ["/cristinaITdeveloper.testchange.dex.MsgSendCreatePair", MsgSendCreatePair],
    ["/cristinaITdeveloper.testchange.dex.MsgCancelBuyOrder", MsgCancelBuyOrder],
];
export const MissingWalletError = new Error("wallet is required");
const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgSendSellOrder: (data) => ({ typeUrl: "/cristinaITdeveloper.testchange.dex.MsgSendSellOrder", value: data }),
        msgCancelSellOrder: (data) => ({ typeUrl: "/cristinaITdeveloper.testchange.dex.MsgCancelSellOrder", value: data }),
        msgSendCreatePair: (data) => ({ typeUrl: "/cristinaITdeveloper.testchange.dex.MsgSendCreatePair", value: data }),
        msgCancelBuyOrder: (data) => ({ typeUrl: "/cristinaITdeveloper.testchange.dex.MsgCancelBuyOrder", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
