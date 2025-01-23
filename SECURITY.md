# Security Policy

## Reporting a Vulnerability

The Foundations team takes security vulnerabilities very seriously. We appreciate your efforts to responsibly disclose your findings and will make every effort to acknowledge your contributions.

### Reporting Process

1. **DO NOT** file a public issue
2. Email your findings to [security@foundations.org](mailto:security@foundations.org)
3. Include as much information as possible:
   - A detailed description of the vulnerability
   - Steps to reproduce the issue
   - Potential impact of the vulnerability
   - Any possible mitigations
   - Your name/handle for credit (optional)

### What to Expect

- You'll receive an acknowledgment within 24 hours
- We'll provide a detailed response within 72 hours
- We'll keep you informed about our progress
- Once fixed, we'll notify you and discuss public disclosure

## Scope

### In Scope

- Nexus (`nexusd`)
- Smart Contract vulnerabilities
- Consensus vulnerabilities
- Network security issues
- Cryptographic weaknesses
- RPC endpoint vulnerabilities

### Out of Scope

- DoS attacks requiring large amounts of traffic
- Issues in dependencies that are already reported
- Social engineering attacks
- Issues requiring physical access
- Issues in deployed contracts not officially associated with Nexus

## Bug Bounty Program

We maintain a bug bounty program for critical security vulnerabilities. Rewards are based on:

- Severity of the vulnerability
- Quality of the report
- Potential impact on the network

### Severity Levels

1. **Critical**: Immediate threat to network security
   - Consensus failures
   - Remote code execution
   - Fund theft

2. **High**: Significant vulnerability
   - Denial of service affecting consensus
   - Significant fund locking

3. **Medium**: Limited impact
   - Non-critical information disclosure
   - Minor fund locking

4. **Low**: Minimal impact
   - Minor optimizations
   - UI/UX issues

## Disclosure Policy

- We follow a coordinated disclosure process
- Vulnerabilities will be disclosed after fixes are available
- Reporters will be credited (if desired)
- Public disclosure timing will be coordinated with the reporter

## Security Best Practices

For validators and node operators:

1. Keep systems updated
2. Use secure key management
3. Follow network upgrade procedures
4. Monitor system resources
5. Implement proper firewall rules
6. Use secure RPC configurations

## Hall of Fame

We maintain a hall of fame for security researchers who have helped improve Nexus's security. Researchers will be credited after the vulnerability has been fixed and disclosed.

## PGP Key

For encrypted communication, please use our PGP key:

```
[PGP KEY TO BE ADDED]
```

## Updates

This security policy may be updated from time to time. Please refer to the git history for changelog.
