import { useEffect, useState } from 'react';
import './App.css';
import axios from 'axios';

interface Transaction {
  Hash: string;
  BlockHeight: string;
  From: string;
  To: string;
  Value: string;
  TxFee: string;
}

export const transaction = (hash: string) => {
  const instance = axios.create({
    baseURL: process.env.REACT_APP_EXPLORER_URL,
  });

  const [transaction, setTransaction] = useState<Transaction>();

  useEffect(() => {
    instance.get(`transaction/hash/${hash}`).then((response) => {
      const data = response.data;
      setTransaction({
        Hash: data.Hash,
        BlockHeight: data.BlockHeight,
        From: data.From,
        To: data.To,
        Value: data.Value,
        TxFee: data.TxFee,
      });
    });
  });

  return (
    <div className="page-transaction">
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
    </div>
  );
};
