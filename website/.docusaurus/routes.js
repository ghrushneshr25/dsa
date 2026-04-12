import React from 'react';
import ComponentCreator from '@docusaurus/ComponentCreator';

export default [
  {
    path: '/dsa/markdown-page',
    component: ComponentCreator('/dsa/markdown-page', '96f'),
    exact: true
  },
  {
    path: '/dsa/',
    component: ComponentCreator('/dsa/', 'f7c'),
    exact: true
  },
  {
    path: '/dsa/',
    component: ComponentCreator('/dsa/', 'bcf'),
    routes: [
      {
        path: '/dsa/',
        component: ComponentCreator('/dsa/', '06c'),
        routes: [
          {
            path: '/dsa/',
            component: ComponentCreator('/dsa/', 'b15'),
            routes: [
              {
                path: '/dsa/recursion-backtracking',
                component: ComponentCreator('/dsa/recursion-backtracking', 'e35'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/dsa/recursion-backtracking/check-if-array-is-sorted-using-recursion',
                component: ComponentCreator('/dsa/recursion-backtracking/check-if-array-is-sorted-using-recursion', '966'),
                exact: true,
                sidebar: "tutorialSidebar"
              },
              {
                path: '/dsa/recursion-backtracking/towers-of-hanoi',
                component: ComponentCreator('/dsa/recursion-backtracking/towers-of-hanoi', '922'),
                exact: true,
                sidebar: "tutorialSidebar"
              }
            ]
          }
        ]
      }
    ]
  },
  {
    path: '*',
    component: ComponentCreator('*'),
  },
];
