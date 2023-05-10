import { useEffect, useState } from 'react';
import axios from 'axios';
import { transaction } from './transaction';

interface Block {
  blockHeight: string;
  receipient: string;
  reward: string;
  size: string;
  gasUsed: string;
  hash: string;
}

interface Transaction {
  Hash: string;
  BlockHeight: string;
  From: string;
  To: string;
  Value: string;
  TxFee: string;
}

export const Main = () => {
  const instance = axios.create({
    baseURL: process.env.REACT_APP_EXPLORER_URL,
  });

  instance.defaults.withCredentials = true;

  const [blocks, setBlocks] = useState<Block[]>([]);
  const [transactions, setTransactions] = useState<Transaction[]>([]);

  useEffect(() => {
    instance.get(`/blocks`).then((response) => {
      const data = response.data;
      console.log(data);
      setBlocks(data);
    });
    instance.get(`/transactions`).then((response) => {
      const data = response.data;
      console.log(data);
      setTransactions(data);
    });
  }, []);

  return (
    <div>
      <div className="input"></div>
      <div className="list-block">
        {blocks.map((block) => (
          <dl>
            <dt>Height</dt>
            <dd>{block?.blockHeight}</dd>
            <dt>Receipient</dt>
            <dd>{block?.receipient}</dd>
            <dt>Reward</dt>
            <dd>{block?.reward}</dd>
            <dt>Size</dt>
            <dd>{block?.size}</dd>
            <dt>GasUsed</dt>
            <dd>{block?.gasUsed}</dd>
            <dt>Hash</dt>
            <dd>{block?.hash}</dd>
          </dl>
        ))}
      </div>
      <div className="list-transaction">
        {transactions?.map((transaction) => (
          <dl>
            <dt>Hash</dt>
            <dd>{transaction?.Hash}</dd>
            <dt>Height</dt>
            <dd>{transaction?.BlockHeight}</dd>
            <dt>From</dt>
            <dd>{transaction?.From}</dd>
            <dt>To</dt>
            <dd>{transaction?.To}</dd>
            <dt>Value</dt>
            <dd>{transaction?.Value}</dd>
            <dt>TxFee</dt>
            <dd>{transaction?.TxFee}</dd>
          </dl>
        ))}
      </div>
    </div>
  );
};
