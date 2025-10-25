// src/components/PaperList.js
import React, { useEffect, useState } from 'react';
import axios from 'axios';

function PaperList() {
  const [papers, setPapers] = useState([]);

  useEffect(() => {
    axios.get('http://localhost:3001/api/papers')
      .then(res => setPapers(res.data))
      .catch(err => console.error(err));
  }, []);

  return (
    <div>
      <h2>期刊论文列表</h2>
      <table border="1" cellPadding="5">
        <thead>
          <tr>
            <th>标题</th><th>作者</th><th>期刊</th><th>发表时间</th>
          </tr>
        </thead>
        <tbody>
          {papers.map(p => (
            <tr key={p.id}>
              <td>{p.title}</td>
              <td>{p.authors}</td>
              <td>{p.journal}</td>
              <td>{p.publish_date}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
export default PaperList;
